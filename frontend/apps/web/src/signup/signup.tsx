function SignUp() {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-4xl font-bold pb-16">Sign Up</h1>
      <div className="flex flex-col items-center justify-center gap-4 w-full max-w-md">
        <input type="text" placeholder="Name" />
        <input type="text" placeholder="Email" />
        <input type="text" placeholder="Username" />
        <input type="password" placeholder="Password" />
        <input type="text" placeholder="Language Preference" />
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
        >
          Sign up
        </button>
      </div>
    </div>
  );
}

export default SignUp;
