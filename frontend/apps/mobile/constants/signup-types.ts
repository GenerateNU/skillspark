export type SignupFormData = {
	language_preference: string;
	email: string;
	password: string;
	confirm_password: string;
	first_name: string;
	last_name: string;
	username: string;
	phone_number: string;
	profile_picture_s3_key: string | undefined;
};

export const signupFormDefaultValues: SignupFormData = {
	language_preference: "",
	email: "",
	password: "",
	confirm_password: "",
	first_name: "",
	last_name: "",
	username: "",
	phone_number: "",
	profile_picture_s3_key: undefined,
};
