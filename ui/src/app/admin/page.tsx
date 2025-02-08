import Users from "./users";
export default async function Admin() {
    const response = await fetch("http://localhost:8000/v1/users");
    return <Users />;
}
