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
  return (
    <Controller
      control={control}
      name={name}
      render={({ field: { onChange, value } }) => (
        <View style={{ width: "100%", gap: 4 }}>
          <TextInput
            style={{
              width: "100%",
              borderWidth: 1,
              borderColor: "#d1d5db",
              borderRadius: 8,
              padding: 10,
              fontSize: 16,
            }}
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
