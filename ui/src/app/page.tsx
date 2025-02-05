"use client";

// useEffect(() => {
//     // useEffect it will run at least once when component mounts
//     // code we want to run
//     // will be runned at least once on component mount
//     return () => {
//         console.log("I am being cleaned up.");
//     }; // optional cleanup function
// }, []); // The dependency array whenever sth changes
// // useEffect will destroy itself and will run this cleanup function before
// // doing so
// // and will be reacreated with new value.

// CSS

import { Button } from "@/components/ui/button";
import { RegisterForm } from "./features/register";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Sun, Moon } from "lucide-react";
import { useTheme } from "next-themes";

export function ModeToggle() {
    const { setTheme } = useTheme();

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="outline" size="icon">
                    <Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
                    <Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
                    <span className="sr-only">Toggle theme</span>
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
                <DropdownMenuItem onClick={() => setTheme("light")}>
                    Light
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => setTheme("dark")}>
                    Dark
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => setTheme("system")}>
                    System
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
export default function Home() {
    return (
        <>
            <div className="p-3">{ModeToggle()}</div>
            <div className="flex justify-center items-center h-screen">
                <Card>
                    <CardHeader>
                        <CardTitle>Welcome to Ukasz App!</CardTitle>
                        <CardDescription>
                            This app is about nothing.
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <Tabs defaultValue="login" className="w-[400px]">
                            <TabsList>
                                <TabsTrigger value="register">
                                    Register
                                </TabsTrigger>
                                <TabsTrigger value="password">
                                    Login
                                </TabsTrigger>
                            </TabsList>
                            <TabsContent value="login"></TabsContent>
                            <TabsContent value="register">
                                <RegisterForm />
                            </TabsContent>
                        </Tabs>
                    </CardContent>
                </Card>
            </div>
        </>
    );
}
