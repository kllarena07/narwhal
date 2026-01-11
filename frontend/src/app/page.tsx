"use client";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
    SelectGroup,
    SelectLabel
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";

export default function Home() {
    async function onSubmit(formData: FormData) {
        try {
            console.log(formData);

            const response = await fetch("http://127.0.0.1:8000/run", {
                method: "POST",
                body: formData
            });

            const data = await response.json();
            console.log(data);
        } catch (err) {
            console.error(err);
        }
    }

    return (
        <main className="flex justify-center items-center w-dvh h-dvh">
            <form className="flex flex-col gap-3 w-[400px]" action={onSubmit}>
                <h1 className="text-xl font-bold">Deploy to Narwhal</h1>
                <Input type="text" placeholder="Link to repository" name="repo" required />
                <Select name="runtime" required>
                    <SelectTrigger className="w-full">
                        <SelectValue placeholder="Select a runtime" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            <SelectLabel>Select a runtime</SelectLabel>
                            <SelectItem value="node">Node</SelectItem>
                        </SelectGroup>
                    </SelectContent>
                </Select>
                <Button className="w-full" type="submit">Create application</Button>
            </form>
        </main>
    );
}
