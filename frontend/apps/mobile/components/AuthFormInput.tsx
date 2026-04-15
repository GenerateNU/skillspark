import { Colors } from "@/constants/theme";
import { Control, Controller, FieldValues, Path } from "react-hook-form";
import { TextInput, TextInputProps, View } from "react-native";

interface AuthFormInputProps<T extends FieldValues> extends TextInputProps {
	control: Control<T>;
	name: Path<T>;
}

export const AuthFormInput = <T extends FieldValues>({
	control,
	name,
	...props
}: AuthFormInputProps<T>) => {
	const colors = Colors["light"];

	return (
		<Controller
			control={control}
			name={name}
			render={({ field: { onChange, value } }) => (
				<View className="w-full">
					<TextInput
						className="w-full rounded-2xl px-4 h-[54px] border-[1px] border-black"
						style={{
							backgroundColor: colors.dropdownBg,
						}}
						autoCorrect={false}
						onChangeText={onChange}
						value={value}
						{...props}
					/>
				</View>
			)}
		/>
	);
};
