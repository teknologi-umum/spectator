export interface User {
    uid: number;
    gid: number;
    free: boolean;
    username: string;
}

export class SystemUsers {
    users: User[];
    constructor(start: number, stop: number, defaultGroup: number) {
        this.users = [];
        for (let i = start; i <= stop; i++) {
            this.users.push({
                uid: i,
                gid: defaultGroup,
                free: false,
                username: `code_executor_${i}`
            });
        }
    }

    public acquire(): User | null {
        const userIndex = this.users.findIndex(u => u.free);
        if (userIndex !== -1) {
            const user = this.users[userIndex];
            this.users[userIndex].free = false;
            return user;
        }
        return null;
    }

    public release(uid: number): void {
        const userIndex = this.users.findIndex(u => u.uid === uid);
        if (userIndex !== -1) {
            this.users[userIndex].free = true;
        }
    }
}
