import { useState } from "react";
import { useForm } from "react-hook-form";
import { useAuthContext } from "../../contexts/use-auth-context";
import { useNavigate } from "react-router-dom";

type SignupFormData = {
	name: string;
	email: string;
	username: string;
	password: string;
	language_preference: string;
	profile_picture_s3_key: string | undefined;
};

function SignUp() {
	const [errorText, setErrorText] = useState("");
	const { signup } = useAuthContext();
	const navigate = useNavigate();

	const { register, handleSubmit } = useForm<SignupFormData>({
		defaultValues: {
			name: "",
			email: "",
			username: "",
			password: "",
			language_preference: "",
			profile_picture_s3_key: undefined,
		},
	});

	const onSubmit = (formData: SignupFormData) => {
		if (
			formData.name === "" ||
			formData.email === "" ||
			formData.username === "" ||
			formData.password === "" ||
			formData.language_preference === ""
		) {
			setErrorText("Missing a required field");
		} else {
			signup(
				formData.name,
				formData.email,
				formData.username,
				formData.password,
				formData.language_preference,
				formData.profile_picture_s3_key,
				setErrorText,
			);
		}
	};

	return (
		<div className="flex flex-col items-center justify-center h-screen">
			<h1 className="text-4xl font-bold pb-16">Sign Up</h1>
			<form
				onSubmit={handleSubmit(onSubmit)}
				className="flex flex-col items-center justify-center gap-4 w-full max-w-md"
			>
				<input type="text" placeholder="Name" {...register("name")} />
				<input type="text" placeholder="Email" {...register("email")} />
				<input type="text" placeholder="Username" {...register("username")} />
				<input
					type="password"
					placeholder="Password"
					{...register("password")}
				/>
				<select
					{...register("language_preference")}
					className="w-full border border-gray-300 rounded-md p-2 text-base bg-white"
				>
					<option value="" disabled>
						Select a language...
					</option>
					<option value="en">English</option>
					<option value="th">Thai</option>
				</select>
				{errorText && <p className="text-red-500 text-sm">{errorText}</p>}
				<button
					type="submit"
					className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
				>
					Sign Up
				</button>
				<button
					type="button"
					onClick={() => navigate("/login")}
					className="text-blue-500 hover:underline"
				>
					Already have an account? Log in
				</button>
			</form>
		</div>
	);
}

export default SignUp;
