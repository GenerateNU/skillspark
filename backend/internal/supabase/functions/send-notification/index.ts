import { serve } from "https://deno.land/std@0.168.0/http/server.ts";
import { Resend } from "npm:resend";
import { messagingApi } from "npm:@line/bot-sdk";
import { createClient } from "npm:@supabase/supabase-js";

const resend = new Resend(Deno.env.get("RESEND_API_KEY"));

const { MessagingApiClient } = messagingApi;
const lineClient = new MessagingApiClient({
    channelAccessToken: Deno.env.get("LINE_CHANNEL_ACCESS_TOKEN") ?? "",
});

const supabaseUrl = Deno.env.get("SUPABASE_URL") ?? "";
const supabaseServiceKey = Deno.env.get("SUPABASE_SERVICE_ROLE_KEY") ?? "";
const supabase = createClient(supabaseUrl, supabaseServiceKey);

serve(async (req) => {
    try {
        const payload = await req.json();

        // The payload structure depends on how the webhook sends it.
        // If it's a database webhook, it might look like { type: 'INSERT', record: { ... } }
        // We assume the payload IS the record or contains the record.
        console.log("Received payload:", JSON.stringify(payload));

        const record = payload.record || payload;

        if (!record || !record.id || !record.type || !record.payload) {
            console.error("Invalid notification record:", record);
            return new Response("Invalid record", { status: 400 });
        }

        const { id, type, payload: msgPayload } = record;

        if (type === "email") {
            const { to, subject, message } = msgPayload;
            if (!to || !message) {
                console.error("Missing email fields");
                return new Response("Missing email fields", { status: 400 });
            }


            console.log(`Sending email to ${to}`);
            const data = await resend.emails.send({
                from: "SkillSpark <onboarding@resend.dev>", // Or configured domain
                to: ["bobbypalazzi@gmail.com"],
                subject: subject || "Notification",
                html: `<p>${message}</p>`,
            });
            console.log("Email sent:", data);

        } else if (type === "sms") {
            // Line Messaging
            // Note: Line requires a User ID, not a phone number usually, unless using specialized services.
            // The assumption here is 'to' contains the Line User ID.
            // If the requirement was strictly SMS via Line, Line doesn't support direct SMS to phone numbers easily without Line Notify or Push Message to User ID.
            // We will assume 'to' is the Line User ID.
            const { to, message } = msgPayload;
            if (!to || !message) {
                console.error("Missing SMS fields");
                return new Response("Missing SMS fields", { status: 400 });
            }

            console.log(`Sending Line message to ${to}`);
            await lineClient.pushMessage({
                to: to,
                messages: [{ type: "text", text: message }],
            });
            console.log("Line message sent");
        }

        // Delete from notifications table
        const { error: deleteError } = await supabase
            .from("notification")
            .delete()
            .eq("id", id);

        if (deleteError) {
            console.error("Failed to delete notification:", deleteError);
            return new Response("Failed to delete", { status: 500 });
        }

        return new Response(JSON.stringify({ success: true }), {
            headers: { "Content-Type": "application/json" },
        });

    } catch (error) {
        console.error("Error processing notification:", error);
        return new Response(JSON.stringify({ error: error.message }), {
            status: 500,
            headers: { "Content-Type": "application/json" },
        });
    }
});
