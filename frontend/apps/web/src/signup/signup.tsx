import { Link } from "react-router-dom";
import { useState } from "react";
import { Btn } from "../components/common";
//import { signupManager } from "@skillspark/api-client";

export default function Signup() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  // const handleSignup = async () => {
  //   try {
  //     if (password !== confirmPassword) {
  //       alert("Passwords do not match");
  //       return;
  //     }

  //     const res = await signupManager({
  //       username,
  //       email,
  //       password,
  //     });

  //     console.log("Signup success:", res);
  //   } catch (err) {
  //     console.error("Signup failed:", err);
  //   }
  // };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-16">Sign Up</h1>

      <div className="flex flex-col gap-4 w-full max-w-md">
        <input
          type="text"
          placeholder="Username"
          className="w-full p-2 border rounded-md"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />

        <input
          type="email"
          placeholder="Email"
          className="w-full p-2 border rounded-md"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          className="w-full p-2 border rounded-md"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <input
          type="password"
          placeholder="Confirm Password"
          className="w-full p-2 border rounded-md"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />

        <Btn
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
          //onClick={handleSignup}
        >
          Create Account
        </Btn>

        <p className="text-sm text-center">
          Already have an account?{" "}
          <Link to="/login" className="text-blue-500 hover:underline">
            Login
          </Link>
        </p>
      </div>
    </div>
  );
}