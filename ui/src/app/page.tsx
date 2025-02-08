"use client";
import { RegisterForm } from "./authorization";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { ModeToggle } from "@/components/mode-toggle";

export default function Authorization() {
    return (
        <>
            <div className="p-3 fixed">{ModeToggle()}</div>
            <div className="h-screen flex justify-center items-center">
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
