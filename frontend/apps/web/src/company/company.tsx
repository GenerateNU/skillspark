import { Routes, Route } from "react-router-dom";
import { CompanySidebar } from "@/components/company_sidebar";
import { CompanyProvider } from "./company-context";
import CompanyHome from "./home";
import CompanyEvents from "./events";
import CompanyPayments from "./payments";
import CompanyCustomers from "./customers";

export default function Company() {
	return (
		<CompanyProvider>
			<div className="flex h-screen w-screen overflow-hidden bg-white font-sans">
				<CompanySidebar />
				<Routes>
					<Route path="" element={<CompanyHome />} />
					<Route path="/events" element={<CompanyEvents />} />
					<Route path="/payments" element={<CompanyPayments />} />
					<Route path="/customers" element={<CompanyCustomers />} />
				</Routes>
			</div>
		</CompanyProvider>
	);
}
