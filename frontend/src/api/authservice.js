export const login = async ({ username, password }) => {
    const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/login`, {
        method: "POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
        mode: "cors",
    });

    if (response.status === 204) throw new Error("Server tidak mengembalikan response.");
    if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);
    return response.json();
};
