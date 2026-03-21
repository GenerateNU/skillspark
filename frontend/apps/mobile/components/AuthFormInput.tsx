import { Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme.web";
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
  const colorScheme = useColorScheme();
  const colors = Colors[colorScheme ?? "light"];

  return (
    <Controller
      control={control}
      name={name}
      render={({ field: { onChange, value } }) => (
        <View style={{ width: "100%", gap: 4 }}>
          <TextInput
            className="w-full rounded-lg p-[10px] text-base"
            style={{ borderWidth: 1, borderColor: colors.borderColor }}
            placeholderTextColor="#9ca3af"
            onChangeText={onChange}
            value={value}
            {...props}
          />
        </View>
      )}
    />
  );
};
