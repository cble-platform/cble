export async function Login(username: string, password: string) {
  return fetch(new URL("/api/auth/login", import.meta.env.VITE_API_BASE_URL), {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      password,
    }),
  });
}

export async function Logout() {
  return fetch(new URL("/api/auth/logout", import.meta.env.VITE_API_BASE_URL), {
    method: "DELETE",
    credentials: "include",
  });
}
