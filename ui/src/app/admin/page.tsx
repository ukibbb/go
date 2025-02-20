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
    const { status, data }: { status: number; data: User[] } =
        await response.json();
    if (status !== 200) {
        return <div>Error: kurwo {status}</div>;
    }
    return <UserTable users={data} />;
}
