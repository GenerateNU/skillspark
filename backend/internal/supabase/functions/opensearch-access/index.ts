import "jsr:@supabase/functions-js/edge-runtime.d.ts";

const OPENSEARCH_URL = Deno.env.get("OPENSEARCH_URL")!;
const OPENSEARCH_USER = Deno.env.get("OPENSEARCH_USER")!;
const OPENSEARCH_PASS = Deno.env.get("OPENSEARCH_PASS")!;
const INDEX = "events";

const authHeader = "Basic " + btoa(`${OPENSEARCH_USER}:${OPENSEARCH_PASS}`);

async function upsert(id: string, doc: Record<string, unknown>) {
  const res = await fetch(`${OPENSEARCH_URL}/${INDEX}/_doc/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      "Authorization": authHeader,
    },
    body: JSON.stringify(doc),
  });
  if (!res.ok) {
    throw new Error(`OpenSearch upsert failed (${res.status}): ${await res.text()}`);
  }
}

async function remove(id: string) {
  const res = await fetch(`${OPENSEARCH_URL}/${INDEX}/_doc/${id}`, {
    method: "DELETE",
    headers: { "Authorization": authHeader },
  });
  if (!res.ok && res.status !== 404) {
    throw new Error(`OpenSearch delete failed (${res.status}): ${await res.text()}`);
  }
}

async function checkConnection(): Promise<boolean> {
  const res = await fetch(`${OPENSEARCH_URL}/_cluster/health`, {
    headers: { "Authorization": authHeader },
  });
  return res.ok;
}

Deno.serve(async (req) => {
  // Health/connection check
  if (req.method === "GET") {
    try {
      const healthy = await checkConnection();
      return new Response(
        JSON.stringify({ connected: healthy, index: INDEX }),
        { status: healthy ? 200 : 503, headers: { "Content-Type": "application/json" } },
      );
    } catch (err) {
      return new Response(JSON.stringify({ connected: false, error: String(err) }), {
        status: 503,
        headers: { "Content-Type": "application/json" },
      });
    }
  }

  try {
    const payload = await req.json();
    const { type, record, old_record } = payload;

    switch (type) {
      case "INSERT":
      case "UPDATE":
        await upsert(record.id, {
          id:                   record.id,
          title_en:             record.title_en,
          title_th:             record.title_th,
          description_en:       record.description_en,
          description_th:       record.description_th,
          category:             record.category,
          header_image_s3_key:  record.header_image_s3_key,
          age_range_min:        record.age_range_min,
          age_range_max:        record.age_range_max,
        });
        break;

      case "DELETE":
        await remove(old_record.id);
        break;

      default:
        return new Response(`Unknown event type: ${type}`, { status: 400 });
    }

    return new Response("ok", { status: 200 });
  } catch (err) {
    console.error(err);
    return new Response(String(err), { status: 500 });
  }
});
