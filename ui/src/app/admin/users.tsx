"use client";
import { User } from "./page";

export default function UserTable(users: User[]) {
    const [search, setSearch] = useState("");
    const filteredUsers = users.filter(
        (user) =>
            user.email.includes(search) ||
            user.username.includes(search) ||
            user.role.includes(search),
    );

    return (
        <div className="p-4">
            <Input
                placeholder="Search users..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                className="mb-4"
            />
            <Table>
                <TableHeader>
                    <TableRow>
                        {columns.map((col) => (
                            <TableHead key={col.id || col.accessorKey}>
                                {col.header}
                            </TableHead>
                        ))}
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {filteredUsers.map((user) => (
                        <TableRow key={user.id}>
                            <TableCell>{user.id}</TableCell>
                            <TableCell>{user.email}</TableCell>
                            <TableCell>{user.username}</TableCell>
                            <TableCell>
                                {user.isActive ? "Active" : "Inactive"}
                            </TableCell>
                            <TableCell>{user.role}</TableCell>
                            <TableCell>
                                <Button variant="outline">Edit</Button>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </div>
    );
}
