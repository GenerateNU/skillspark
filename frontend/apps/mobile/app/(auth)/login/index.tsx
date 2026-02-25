import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useState } from "react";
import { Button, StyleSheet, TextInput } from "react-native";
import { useLoginGuardian } from '@skillspark/api-client';

export default function LoginScreen() {
  const [emailText, setEmailText] = useState("");
  const [passwordText, setPasswordText] = useState("");

  const handleLogIn = () => {
    if ((emailText === "") || (passwordText === "")) {
      throw new Error("Missing email or password");
    }

    try {
      useLoginGuardian()
    }
  }
  
  return (
        <ThemedView style={styles.titleContainer}>
          <ThemedText type="title">Log In</ThemedText>
          <TextInput
          placeholder="Email"
          onChangeText={setEmailText}
          value={emailText}
          keyboardType="email-address"
          />
          <TextInput
          placeholder="Password"
          onChangeText={setPasswordText}
          value={passwordText}
          keyboardType="default"
          secureTextEntry={true}
          />
          <Button
            title={"Log In"}
            onPress={handleLogIn()}
          />
        </ThemedView>
     );
}

const styles = StyleSheet.create({
  titleContainer: {
    flexDirection: "row",
    alignItems: "center",
    gap: 8,
  },
  errorText: {
    color: '#ff4444',
  },
});