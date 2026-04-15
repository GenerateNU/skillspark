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
  const colors = Colors.light;

  return (
    <Controller
      control={control}
      name={name}
      render={({ field: { onChange, value } }) => (
        <View className="w-full gap-1">
          <TextInput
            className="w-full border rounded-lg px-[10px] text-base h-12 leading-5"
            style={{
              borderColor: colors.borderColor,
            }}
            placeholderTextColor={AppColors.placeholderText}
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
