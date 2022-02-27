export interface User {
    uid: number;
    gid: number;
    free: boolean;
}

export class SystemUsers {
    users: User[];
    constructor(start: number, stop: number, defaultGroup: number) {
        const users: User[] = [];

        for (let i = start; i <= stop; i++) {
            users.push({
                uid: i,
                gid: defaultGroup,
                free: false
            });
        }
        this.users = users;
    }

    public acquire(): User | null {
        const user = this.users.find(u => u.free);
        if (user) {
            user.free = false;
            this.users = [
                ...this.users.filter(u => u.uid !== user.uid),
                user
            ];
            return user;
        }
        return null;
    }

    public release(uid: number): void {
        const user = this.users.find(u => u.uid === uid);
        if (user) {
            user.free = true;
            this.users = [
                ...this.users.filter(u => u.uid !== user.uid),
                user
            ];
        }
    }
}
