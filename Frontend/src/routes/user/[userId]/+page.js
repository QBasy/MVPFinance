async function getUserByID(id) {
    const res = fetch('http://localhost:8080/api');
    return await res.json();
}