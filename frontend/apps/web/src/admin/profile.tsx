export function ProfilePage() {
  
  const profile = {
    name: import.meta.env.VITE_ADMIN_NAME ?? "—",
    email: import.meta.env.VITE_ADMIN_EMAIL ?? "—",
    username: import.meta.env.VITE_ADMIN_USERNAME ?? "—",
    language_preference: import.meta.env.VITE_ADMIN_LANG_PREFERENCE ?? "en",
    role: import.meta.env.VITE_ADMIN_ROLE ?? "—",
  };

  const languageLabels: Record<string, string> = { en: "English", th: "Thai" };

  const profileFields = [
              { label: "Name", value: profile.name },
              { label: "Email", value: profile.email },
              { label: "Username", value: profile.username },
              { label: "Language", value: languageLabels[profile.language_preference] ?? profile.language_preference },
              { label: "Role", value: profile.role },
            ]

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      <div className="bg-white border-b border-gray-200 px-6 py-4 shrink-0">
        <h1 className="text-base font-semibold text-gray-900">Profile</h1>
      </div>

      <div className="flex-1 overflow-auto bg-gray-50 p-6 flex justify-center">
        <div className="w-full max-w-lg flex flex-col gap-6">

          <div className="bg-white rounded-lg border border-gray-200 px-5 py-6 flex flex-col items-center gap-3 text-center">
            <div className="w-20 h-20 rounded-full bg-blue-600 flex items-center justify-center text-3xl font-bold text-white">
              {profile.name.charAt(0).toUpperCase()}
            </div>
            <div>
              <p className="text-base font-semibold text-gray-900">{profile.name}</p>
              <p className="text-sm text-gray-500">{profile.email}</p>
              <span className="inline-flex items-center text-xs font-medium px-2 py-0.5 rounded ring-1 bg-blue-50 text-blue-700 ring-blue-200 mt-1">
                {profile.role}
              </span>
            </div>
          </div>

          <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
            <div className="px-5 py-4">
              <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">Account details</h2>
            </div>
            {profileFields.map(function (row: { label: string; value: string }) {
              return (
                <div key={row.label} className="px-5 py-3 grid grid-cols-3 gap-4">
                  <span className="text-xs font-medium text-gray-500">{row.label}</span>
                  <span className="col-span-2 text-sm text-gray-800">{row.value}</span>
                </div>
              );
            })}
          </div>

        </div>
      </div>
    </div>
  );
}