function Greeting() {
    return <h1>Hello there</h1>;
}

function Greeting(props: { name: string }) {
    return <h1>Hello {props.name}</h1>;
}

<Greeting name="Mudi" />


import { useState } from "react";

function Counter() {
    const [count, setCount] = useState(0);
    return <button onClick={() => setCount(count + 1)}>{count}</button>
}

function Button({ onClick }: { onClick: () => void }) {
    return <button onClick={onClick}>Click me</button>;
}

<Button onClick={() => console.log("clicked")} />





// Basic Types

let age: number = 27;
let name: string = "Mudi";
let isActive: boolean = true;
let tags: string[] = ["react", "typescript"];


// Custom Types

type Developer = {
    name: string;   
    yearsOfExperienc: number;
    isActive: boolean;
};

const mudi: Developer = {
    name: "Mudi";
    yearsOfExperience: 3;
    isActive: true;
};


// FunctionTypes
function double(value: number): number {
    return value * 2;
}

// Union Types
function formatId(id: string | number): string {
    return `ID-${id}`;
}

// Generics

function wrapInArray<T>(value: T): T[] {
    return [value];
}

wrapInArray<string>("Mudi");
wrapInArray<number>(57);


