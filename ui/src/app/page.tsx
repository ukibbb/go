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

import { RegisterForm } from "./features/register";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

export default function Home() {
    return (
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
                            <TabsTrigger value="register">Register</TabsTrigger>
                            <TabsTrigger value="password">Login</TabsTrigger>
                        </TabsList>
                        <TabsContent value="login"></TabsContent>
                        <TabsContent value="register">
                            <RegisterForm />
                        </TabsContent>
                    </Tabs>
                </CardContent>
            </Card>
        </div>
    );
}
