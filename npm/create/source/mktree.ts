import path from "node:path";
import fs from "node:fs";

export interface FileNode {
    readonly type: "filenode",
    readonly content: Uint8Array,
    readonly name: string,
}

export interface DirNode {
    readonly type: "dirnode",
    readonly nodes: FSNode[],
    readonly name: string,
}

export type FSNode = FileNode | DirNode;

export function createFileNode(name: string, content: string | Uint8Array): FileNode {
    return {
        type: "filenode",
        name,
        content: (() => {
            if (typeof content === "string") {
                return new TextEncoder().encode(content.replaceAll("\n", "\r\n"));
            }
            return content;
        })(),
    }
}

export function createDirNode(name: string, nodes: FSNode[]): DirNode {
    return {
        type: "dirnode",
        name,
        nodes,
    }
}

export function writeFS(root: FSNode) {
    writeFSInternal(".", root);
}

function writeFSInternal(dirname: string, node: FSNode) {
    if (node.type === "filenode") {
        const filePath = path.join(dirname, node.name);
        fs.writeFileSync(filePath, node.content, "binary");
    } else {
        for (const childNode of node.nodes) {
            if (childNode.type === "dirnode") {
                const nodePath = path.join(dirname, childNode.name);
                fs.mkdirSync(nodePath, { recursive: true })
                writeFSInternal(nodePath, childNode);
            } else {
                writeFSInternal(dirname, childNode);
            }
        }
    }
}
