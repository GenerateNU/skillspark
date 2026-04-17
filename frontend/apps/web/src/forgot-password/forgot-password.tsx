import { useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useForgotPassword } from "@skillspark/api-client";

type ForgotPasswordFormData = {
  email: string;
};

export default function ForgotPassword() {
  const [errorText, setErrorText] = useState("");
  const [submitted, setSubmitted] = useState(false);
  const navigate = useNavigate();
  const { mutate: forgotPasswordFunc, isPending } = useForgotPassword();

  const { register, handleSubmit } = useForm<ForgotPasswordFormData>({
    defaultValues: { email: "" },
  });

  const onSubmit = (formData: ForgotPasswordFormData) => {
    if (!formData.email) {
      setErrorText("Please enter your email address");
      return;
    }
    setErrorText("");
    forgotPasswordFunc(
      { data: { email: formData.email } },
      {
        onSuccess: () => setSubmitted(true),
        onError: () => {
          // Always show the generic message to avoid leaking email existence
          setSubmitted(true);
        },
      },
    );
  };

  if (submitted) {
    return (
      <div className="flex flex-col items-center justify-center h-screen">
        <h1 className="text-4xl font-bold pb-4">Check Your Email</h1>
        <p className="text-gray-600 text-center max-w-sm pb-8">
          If that email is registered, you'll receive password reset
          instructions shortly.
        </p>
        <button
          type="button"
          onClick={() => navigate("/login")}
          className="text-blue-500 hover:underline"
        >
          Back to Log In
        </button>
      </div>
    );
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-4">Forgot Password</h1>
      <p className="text-gray-600 text-center max-w-sm pb-8">
        Enter your email address and we'll send you instructions to reset your
        password.
      </p>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col items-center justify-center gap-4 w-full max-w-md"
      >
        <input
          type="email"
          placeholder="Email"
          {...register("email")}
          className="w-full"
        />
        {errorText && <p className="text-red-500 text-sm">{errorText}</p>}
        <button
          type="submit"
          disabled={isPending}
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 disabled:opacity-50"
        >
          {isPending ? "Sending..." : "Send Reset Link"}
        </button>
        <button
          type="button"
          onClick={() => navigate("/login")}
          className="text-blue-500 hover:underline"
        >
          Back to Log In
        </button>
      </form>
    </div>
  );
}
