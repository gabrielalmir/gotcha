export class Base62Converter {
    private static BASE62 = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';

    public static toBase62(value: number): string {
        if (value === 0) {
            return '0';
        }

        let sb = '';
        while (value > 0) {
            const remainder = value % 62;
            sb = Base62Converter.BASE62.charAt(remainder) + sb;
            value = Math.floor(value / 62);
        }
        return sb;
    }
}
