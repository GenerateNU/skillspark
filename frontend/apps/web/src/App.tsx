import { useAuthContext } from "./contexts/use-auth-context";

function App() {
	const { logout } = useAuthContext();

	return (
		<div className="w-screen p-8 flex justify-center items-center">
			<div className="w-screen p-8 flex justify-center items-center">
				<h1 className="text-4xl font-bold">Welcome to SkillSpark!</h1>
			</div>
			<div className="absolute bottom-0 justify-center items-center">
				<button
					type="submit"
					className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
					onClick={() => logout()}
				>
					Log Out
				</button>
			</div>
		</div>
	);
}

export default App;
