import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import React, { useState } from "react";
import { TextInput, TouchableOpacity, Text, View } from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useForm, Controller } from "react-hook-form";

type LoginFormData = {
  email: string;
  password: string;
};

export default function LoginScreen() {
  const [errorText, setErrorText] = useState("");
  const { login } = useAuthContext();

  const { control, handleSubmit } = useForm<LoginFormData>({
    defaultValues: {
      email: "",
      password: "",
    }
  });

  const onSubmit = (formData: LoginFormData) => {
    if (formData.email === "" || formData.password === "") {
      setErrorText("Missing email or password")
    } else {
      login(formData.email, formData.password, setErrorText);
    }
  };

  const handleGoToSignUp = () => {
    router.push("/(auth)/signup");
  };

  const inputStyle = {
    width: "100%" as const,
    borderWidth: 1,
    borderColor: "#d1d5db",
    borderRadius: 8,
    padding: 10,
    fontSize: 16,
  };

  return (
    <ThemedView style={{ flex: 1, alignItems: "center", justifyContent: "center" }}>
      <ThemedText type="title" style={{ fontSize: 30, fontWeight: "bold", marginBottom: 30 }}>
        Log In
      </ThemedText>

      <View style={{ width: "100%", paddingHorizontal: 24, gap: 16, alignItems: "center" }}>
        <Controller
          control={control}
          name="email"
          render={({ field: { onChange, value } }) => (
            <View style={{ width: "100%", gap: 4 }}>
              <TextInput
                style={inputStyle}
                placeholder="Email"
                onChangeText={onChange}
                value={value}
                keyboardType="email-address"
                autoCapitalize="none"
              />
            </View>
          )}
        />
        <Controller
          control={control}
          name="password"
          render={({ field: { onChange, value } }) => (
            <View style={{ width: "100%", gap: 4 }}>
              <TextInput
                style={inputStyle}
                placeholder="Password"
                onChangeText={onChange}
                value={value}
                secureTextEntry={true}
              />
            </View>
          )}
        />
        <TouchableOpacity
          style={{
            backgroundColor: "#3b82f6",
            borderRadius: 8,
            padding: 10,
            width: "100%",
            alignItems: "center",
          }}
          onPress={handleSubmit(onSubmit)}
          activeOpacity={0.5}
        >
          <Text style={{ color: "white", fontSize: 16, fontWeight: "500" }}>Log In</Text>
        </TouchableOpacity>
        <TouchableOpacity
          style={{
            borderRadius: 8,
            padding: 10,
            width: "100%",
            alignItems: "center",
          }}
          onPress={handleGoToSignUp}
          activeOpacity={0.5}
        >
        <Text style={{ color: "#3b82f6", fontSize: 16, fontWeight: "500" }}>Don&apos;t have an account? Sign up</Text>
        </TouchableOpacity>
        <Text style={{ color: "#ef4444", fontSize: 16, textAlign: "center" }}>{errorText}</Text>
      </View>
    </ThemedView>
  );
}
