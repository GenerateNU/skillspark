import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import { useState } from "react";
import { TextInput, TouchableOpacity, Text, View } from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";

export default function SignupScreen() {
  const [nameText, setNameText] = useState("");
  const [emailText, setEmailText] = useState("");
  const [usernameText, setUsernameText] = useState("");
  const [passwordText, setPasswordText] = useState("");
  const [langPrefText, setLangPrefText] = useState("");
  const [s3KeyText] = useState(undefined);
  const [errorText, setErrorText] = useState("");
  const { signup } = useAuthContext();

  const handleSignUp = () => {
    if (nameText === "" 
      || emailText === "" 
      || usernameText === ""
      || passwordText === ""
      || langPrefText === "") {
      setErrorText("Missing a required field");
    } else {
      signup(nameText, emailText, usernameText, passwordText, langPrefText, s3KeyText, setErrorText); 
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
        <TextInput 
          style={inputStyle} 
          placeholder="Full Name" 
          onChangeText={setNameText} 
          value={nameText} 
          autoCapitalize="none" 
        />
        <TextInput 
          style={inputStyle} 
          placeholder="Email" 
          onChangeText={setEmailText} 
          value={emailText} 
          keyboardType="email-address" 
          autoCapitalize="none" 
        />
        <TextInput 
          style={inputStyle} 
          placeholder="Username" 
          onChangeText={setUsernameText} 
          value={usernameText} 
          autoCapitalize="none" 
        />
        <TextInput 
          style={inputStyle} 
          placeholder="Password" 
          onChangeText={setPasswordText} 
          value={passwordText} 
          secureTextEntry={true} 
        />
        <TextInput 
          style={inputStyle} 
          placeholder="Language Preference" 
          onChangeText={setLangPrefText} 
          value={langPrefText} 
          autoCapitalize="none" 
        />
        <TouchableOpacity
          style={{
            backgroundColor: "#3b82f6",
            borderRadius: 8,
            padding: 10,
            width: "100%",
            alignItems: "center",
          }}
          onPress={handleSignUp}
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