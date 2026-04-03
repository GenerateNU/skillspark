import { AppColors, Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme.web";
import {
	Control,
	Controller,
	FieldError,
	FieldValues,
	Path,
} from "react-hook-form";
import { Text, TextInput, TextInputProps, View } from "react-native";

interface AuthFormInputProps<T extends FieldValues> extends TextInputProps {
	control: Control<T>;
	name: Path<T>;
	error: FieldError | undefined;
}

export const AuthFormInput = <T extends FieldValues>({
	control,
	name,
	error,
	...props
}: AuthFormInputProps<T>) => {
	const colorScheme = useColorScheme();
	const colors = Colors[colorScheme ?? "light"];

	return (
		<Controller
			control={control}
			name={name}
			rules={{ required: name + " is required" }}
			render={({ field: { onChange, value } }) => (
				<View className="w-full gap-1">
					<TextInput
						className="w-full border rounded-lg px-[10px] py-[10px] text-base"
						style={{ borderColor: colors.borderColor }}
						placeholderTextColor={AppColors.placeholderText}
						onChangeText={onChange}
						value={value}
						{...props}
					/>
					{error && (
						<Text style={{ color: AppColors.danger, fontSize: 13 }}>
							{error.message}
						</Text>
					)}
				</View>
			)}
		/>
	);
};
