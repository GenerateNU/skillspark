import { useForm } from "react-hook-form";
import { useState } from "react";
import { useAuthContext } from "../../contexts/use-auth-context";
import { useNavigate } from "react-router-dom";

type LoginFormData = {
  email: string;
  password: string;
};

export default function Login() {
  const [errorText, setErrorText] = useState("");
  const { login } = useAuthContext();
  const navigate = useNavigate();
  
  const { register, handleSubmit } = useForm<LoginFormData>({
    defaultValues: {
      email: "",
      password: "",
    }
  });

  const onSubmit = (formData: LoginFormData) => {
    if (formData.email === "" || formData.password === "") {
      setErrorText("Missing email or password");
    } else {
      login(formData.email, formData.password, setErrorText);
    }
  };

    return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-16">Log In</h1>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col items-center justify-center gap-4 w-full max-w-md"
      >
        <input
          type="text"
          placeholder="Email"
          {...register("email")}
        />
        <input
          type="password"
          placeholder="Password"
          {...register("password")}
        />
        {errorText && <p className="text-red-500 text-sm">{errorText}</p>}
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
        >
          Log In
        </button>
        <button
          type="button"
          onClick={() => navigate("/signup")}
          className="text-blue-500 hover:underline"
        >
          Don't have an account? Sign up
        </button>
      </form>
    </div>
  );
}