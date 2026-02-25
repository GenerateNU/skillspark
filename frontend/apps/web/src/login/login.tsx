import { useState, type SetStateAction } from "react";

export default function Login() {
  const [emailText, setEmailText] = useState("");
  const [passwordText, setPasswordText] = useState("");

  const handleEmailChange = (event: { target: { value: SetStateAction<string>; }; }) => {
    setEmailText(event.target.value);
  };

  const handlePasswordChange = (event: { target: { value: SetStateAction<string>; }; }) => {
    setPasswordText(event.target.value);
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-16">Log In</h1>
      <div className="flex flex-col items-center justify-center gap-4 w-full max-w-md">
        <input 
        type="text" 
        placeholder="Email" 
        onChange={handleEmailChange}
        value={emailText}
        />
        <input 
        type="password" 
        placeholder="Password" 
        onChange={handlePasswordChange}
        value={passwordText}
        />
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
        >
          Log in
        </button>
      </div>
    </div>
  );
}

