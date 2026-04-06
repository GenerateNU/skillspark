import { Navigate } from "react-router-dom";
import { useAuthContext } from "./contexts/use-auth-context";

export default function ProtectedRoute({
	children,
}: {
	children: React.ReactNode;
}) {
	const { isAuthenticated } = useAuthContext();

	if (!isAuthenticated) {
		return <Navigate to="/login" />;
	}

	return <>{children}</>;
}
