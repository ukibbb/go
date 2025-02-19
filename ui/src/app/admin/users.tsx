"use client";
import React from "react";
import { User } from "./page";
import { Input } from "@/components/ui/input";
import {
    Table,
    TableRow,
    TableCell,
    TableHead,
    TableBody,
    TableHeader,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";

import { ColumnDef } from "@tanstack/react-table";

const columns: ColumnDef<User>[] = [
    {
        accessorKey: "id",
        header: "ID",
    },
    {
        accessorKey: "email",
        header: "Email",
    },
    {
        accessorKey: "username",
        header: "Username",
    },
    {
        accessorKey: "isActive",
        header: "Status",
        cell: ({ row }) => (row.original.isActive ? "Active" : "Inactive"),
    },
    {
        accessorKey: "role",
        header: "Role",
    },
    {
        id: "actions",
        header: "Actions",
        cell: ({ row }) => <Button variant="outline">Edit</Button>,
    },
];

interface UserTableProps {
    users: User[];
}

export default function UserTable(props: UserTableProps) {
    const { users } = props;
    const [search, setSearch] = React.useState("");
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
