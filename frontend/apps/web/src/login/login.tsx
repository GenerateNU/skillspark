export default function Login() {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-16">Login</h1>
      <div className="flex flex-col items-center justify-center gap-4 w-full max-w-md">
        <input type="text" placeholder="Username" />
        <input type="password" placeholder="Password" />
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
        >
          Login
        </button>
      </div>
    </div>
  );
}
