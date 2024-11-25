"use client";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
    SelectGroup
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
        <main>
            <h1>Welcome to Narwhal</h1>
            <form className="flex flex-col" action={onSubmit}>
                <Label>
                    <Input type="text" placeholder="Link to repository" name="repo" />
                </Label>
                <Select name="runtime">
                    <SelectTrigger className="w-[180px]">
                        <SelectValue placeholder="Select a runtime" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            <SelectItem value="node">Node</SelectItem>
                        </SelectGroup>
                    </SelectContent>
                </Select>
                <Button type="submit">Create application</Button>
            </form>
        </main>
    );
}
