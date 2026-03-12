import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import { useState } from "react";
import { TextInput, TouchableOpacity, Text, View } from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";
import { Controller, useForm } from "react-hook-form";

type SignupFormData = {
  name: string;
  email: string;
  username: string;
  password: string;
  language_preference: string;
  profile_picture_s3_key: string | undefined;
};

export default function SignupScreen() {
  const [errorText, setErrorText] = useState("");
  const { signup } = useAuthContext();

  const { control, handleSubmit } = useForm<SignupFormData>({
      defaultValues: {
        name: "",
        email: "",
        username: "",
        password: "",
        language_preference: "",
        profile_picture_s3_key: undefined
      }
    }
  );

  const onSubmit = (formData: SignupFormData) => {
    if (formData.name === "" 
      || formData.email === "" 
      || formData.username === ""
      || formData.password === ""
      || formData.language_preference === "") {
      setErrorText("Missing a required field");
    } else {
      signup(
        formData.name, 
        formData.email, 
        formData.username, 
        formData.password, 
        formData.language_preference, 
        formData.profile_picture_s3_key, 
        setErrorText
      ); 
    }
  };

  const handleGoToLogIn = () => {
    router.push("/(auth)/login");
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
        Sign Up
      </ThemedText>
      <View style={{ width: "100%", paddingHorizontal: 24, gap: 16, alignItems: "center" }}>
        <Controller
        control={control}
        name="name"
        render={({ field: { onChange, value } }) => (
          <View style={{ width: "100%", gap: 4 }}>
            <TextInput
              style={inputStyle}
              placeholder="Full Name"
              onChangeText={onChange}
              value={value}
              autoCapitalize="none"
            />
          </View>
        )}
      />
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
        name="username"
        render={({ field: { onChange, value } }) => (
          <View style={{ width: "100%", gap: 4 }}>
            <TextInput
              style={inputStyle}
              placeholder="Username"
              onChangeText={onChange}
              value={value}
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
      <Controller
        control={control}
        name="language_preference"
        render={({ field: { onChange, value } }) => (
          <View style={{ width: "100%", gap: 4 }}>
            <TextInput
              style={inputStyle}
              placeholder="Language Preference"
              onChangeText={onChange}
              value={value}
              autoCapitalize="none"
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
          activeOpacity={0.8}
        >
          <Text style={{ color: "white", fontSize: 16, fontWeight: "500" }}>Sign Up</Text>
        </TouchableOpacity>
        <TouchableOpacity
          style={{
            borderRadius: 8,
            padding: 10,
            width: "100%",
            alignItems: "center",
          }}
          onPress={handleGoToLogIn}
          activeOpacity={0.5}
        >
        <Text style={{ color: "#3b82f6", fontSize: 16, fontWeight: "500" }}>Already have an account? Log in</Text>
        </TouchableOpacity>
        <Text style={{ color: "#ef4444", fontSize: 16, textAlign: "center" }}>{errorText}</Text>
      </View>
    </ThemedView>
  );
}