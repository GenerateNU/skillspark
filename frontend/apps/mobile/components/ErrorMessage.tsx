import { Text } from "react-native";

interface ErrorMessageProps {
  message: string;
}

export const ErrorMessage = ({ message }: ErrorMessageProps) => {
  return (
    <Text style={{ color: "#ef4444", fontSize: 16, textAlign: "center" }}>
      {message}
    </Text>
  );
};
