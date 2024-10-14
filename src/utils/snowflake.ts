export class Snowflake {
    private static readonly EPOCH = 1288834974657;
    private static readonly WORKER_ID_BITS = 5;
    private static readonly DATACENTER_ID_BITS = 5;
    private static readonly SEQUENCE_BITS = 12;

    private static readonly MAX_WORKER_ID = ~(-1 << Snowflake.WORKER_ID_BITS);
    private static readonly MAX_DATACENTER_ID = ~(-1 << Snowflake.DATACENTER_ID_BITS);
    private static readonly SEQUENCE_MASK = ~(-1 << Snowflake.SEQUENCE_BITS);

    private static readonly WORKER_ID_SHIFT = Snowflake.SEQUENCE_BITS;
    private static readonly DATACENTER_ID_SHIFT = Snowflake.SEQUENCE_BITS + Snowflake.WORKER_ID_BITS;
    private static readonly TIMESTAMP_LEFT_SHIFT = Snowflake.SEQUENCE_BITS + Snowflake.WORKER_ID_BITS + Snowflake.DATACENTER_ID_BITS;

    private workerId: number;
    private datacenterId: number;
    private sequence: number = 0;
    private lastTimestamp: number = -1;

    constructor(workerId: number, datacenterId: number) {
        if (workerId > Snowflake.MAX_WORKER_ID || workerId < 0) {
            throw new Error(`Worker ID deve estar entre 0 e ${Snowflake.MAX_WORKER_ID}`);
        }
        if (datacenterId > Snowflake.MAX_DATACENTER_ID || datacenterId < 0) {
            throw new Error(`Datacenter ID deve estar entre 0 e ${Snowflake.MAX_DATACENTER_ID}`);
        }
        this.workerId = workerId;
        this.datacenterId = datacenterId;
    }

    public nextId(): number {
        let timestamp = this.currentTime();

        if (timestamp < this.lastTimestamp) {
            throw new Error("Clock está indo para trás. Recusando gerar ID.");
        }

        if (this.lastTimestamp === timestamp) {
            this.sequence = (this.sequence + 1) & Snowflake.SEQUENCE_MASK;

            if (this.sequence === 0) {
                timestamp = this.waitNextMillis(this.lastTimestamp);
            }
        } else {
            this.sequence = 0;
        }

        this.lastTimestamp = timestamp;

        return ((timestamp - Snowflake.EPOCH) << Snowflake.TIMESTAMP_LEFT_SHIFT)
            | (this.datacenterId << Snowflake.DATACENTER_ID_SHIFT)
            | (this.workerId << Snowflake.WORKER_ID_SHIFT)
            | this.sequence;
    }

    private waitNextMillis(lastTimestamp: number): number {
        let timestamp = this.currentTime();
        while (timestamp <= lastTimestamp) {
            timestamp = this.currentTime();
        }
        return timestamp;
    }

    private currentTime(): number {
        return Date.now();
    }
}
