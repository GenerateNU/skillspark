import { Link } from "react-router-dom";
import { useState } from "react";
import { Btn } from "../components/common";
import { loginManager } from "../../../../packages/api-client/src/generated/auth/auth";

export default function Login() {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleLogin = async () => {
    try {
      const data = await loginManager({email: username, password})

      console.log("Logged in:", data);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-16">Login</h1>

      <div className="flex flex-col gap-4 w-full max-w-md">
        <input
          type="text"
          placeholder="Username"
          className="w-full p-2 border rounded-md"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          className="w-full p-2 border rounded-md"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <Btn
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
          onClick={handleLogin}
        >
          Login
        </Btn>

        <p className="text-sm text-center">
          Don't have an account?{" "}
          <Link to="/signup" className="text-blue-500 hover:underline">
            Sign up
          </Link>
        </p>
      </div>
    </div>
  );
}

