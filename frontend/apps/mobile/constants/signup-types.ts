export type SignupFormData = {
	language_preference: string;
	email: string;
	password: string;
	confirm_password: string;
	name: string;
	username: string;
	profile_picture_s3_key: string | undefined;
};

export const signupFormDefaultValues: SignupFormData = {
	language_preference: "",
	email: "",
	password: "",
	confirm_password: "",
	name: "",
	username: "",
	profile_picture_s3_key: undefined,
};
