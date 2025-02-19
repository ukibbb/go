import UserTable from "./users";

export type User = {
    id: string;
    email: string;
    username: string;
    isActive: boolean;
    role: string;
};
export default async function Admin() {
    const response = await fetch("http://localhost:8000/v1/users");
    const users: { status: number; msg: User[] } = await response.json();
    return <UserTable users={users.msg} />;
}
