import { AppColors, Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme.web";
import { Control, Controller, FieldValues, Path } from "react-hook-form";
import { TextInput, TextInputProps, View, ScrollView } from "react-native";

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
    <ScrollView
      keyboardShouldPersistTaps="handled"
    >
      <Controller
        control={control}
        name={name}
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
          </View>
        )}
      />
    </ScrollView>
  );
};
