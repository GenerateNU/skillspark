import { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useResetPassword } from "@skillspark/api-client";

type ResetPasswordFormData = {
  newPassword: string;
  confirmPassword: string;
};

const PASSWORD_RULES = [
  { key: "length", check: (p: string) => p.length >= 8, label: "At least 8 characters" },
  { key: "upper", check: (p: string) => /[A-Z]/.test(p), label: "At least one uppercase letter" },
  { key: "lower", check: (p: string) => /[a-z]/.test(p), label: "At least one lowercase letter" },
  { key: "number", check: (p: string) => /[0-9]/.test(p), label: "At least one number" },
  { key: "special", check: (p: string) => /[!@#~$%^&*()+|_.,;<>?/{}\\-]/.test(p), label: "At least one special character" },
];

export default function ResetPassword() {
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const [errorText, setErrorText] = useState("");
  const [success, setSuccess] = useState(false);
  const navigate = useNavigate();
  const { mutate: resetPasswordFunc, isPending } = useResetPassword();

  const { register, handleSubmit, watch } = useForm<ResetPasswordFormData>({
    defaultValues: { newPassword: "", confirmPassword: "" },
  });

  const newPassword = watch("newPassword");

  useEffect(() => {
    // Supabase appends access_token in the URL hash fragment after redirect
    const hash = window.location.hash;
    const params = new URLSearchParams(hash.replace("#", "?"));
    const token = params.get("access_token");
    const type = params.get("type");

    if (token && type === "recovery") {
      setAccessToken(token);
    } else {
      setErrorText("Invalid or expired reset link. Please request a new one.");
    }
  }, []);

  const onSubmit = (formData: ResetPasswordFormData) => {
    if (!accessToken) {
      setErrorText("Invalid or expired reset link. Please request a new one.");
      return;
    }

    const failed = PASSWORD_RULES.filter((r) => !r.check(formData.newPassword));
    if (failed.length > 0) {
      setErrorText(
        "Password does not meet requirements:\n" +
          failed.map((r) => `• ${r.label}`).join("\n"),
      );
      return;
    }

    if (formData.newPassword !== formData.confirmPassword) {
      setErrorText("Passwords do not match");
      return;
    }

    setErrorText("");
    resetPasswordFunc(
      {
        data: {
          access_token: accessToken,
          new_password: formData.newPassword,
        },
      },
      {
        onSuccess: () => setSuccess(true),
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string } };
          setErrorText(fail.data?.message ?? "Failed to reset password. Please try again.");
        },
      },
    );
  };

  if (success) {
    return (
      <div className="flex flex-col items-center justify-center h-screen">
        <h1 className="text-4xl font-bold pb-4">Password Updated</h1>
        <p className="text-gray-600 text-center max-w-sm pb-8">
          Your password has been updated successfully.
        </p>
        <button
          type="button"
          onClick={() => navigate("/login")}
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
        >
          Log In
        </button>
      </div>
    );
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-8">Reset Password</h1>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col items-center justify-center gap-4 w-full max-w-md"
      >
        <div className="w-full">
          <input
            type="password"
            placeholder="New Password"
            {...register("newPassword")}
            className="w-full"
          />
          {newPassword.length > 0 && (
            <ul className="mt-2 space-y-1">
              {PASSWORD_RULES.map((rule) => {
                const met = rule.check(newPassword);
                return (
                  <li
                    key={rule.key}
                    className={`text-xs flex items-center gap-1 ${met ? "text-green-600" : "text-gray-400"}`}
                  >
                    <span>{met ? "✓" : "✗"}</span>
                    {rule.label}
                  </li>
                );
              })}
            </ul>
          )}
        </div>
        <input
          type="password"
          placeholder="Confirm New Password"
          {...register("confirmPassword")}
          className="w-full"
        />
        {errorText && (
          <p className="text-red-500 text-sm whitespace-pre-line">{errorText}</p>
        )}
        <button
          type="submit"
          disabled={isPending || !accessToken}
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 disabled:opacity-50"
        >
          {isPending ? "Updating..." : "Update Password"}
        </button>
      </form>
    </div>
  );
}
