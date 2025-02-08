"use client";
import { ModeToggle } from "@/components/mode-toggle";
export default function AdminLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <>
            <div className="p-3">{ModeToggle()}</div>
            <div className="w-screen h-screen flex justify-center items-center">
                {children}
            </div>
        </>
    );
}
