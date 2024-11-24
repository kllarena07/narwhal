"use client";

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
                <label>
                    Link to repository:
                    <input className="text-black" type="text" name="repo" />
                </label>
                <select className="text-black" name="app_type">
                    <option value="node">Node</option>
                </select>
                <button type="submit">Create application</button>
            </form>
        </main>
    );
}
