var Hg = ["=", ".", "-", "*"], Ig = 8388608;
var Dg = 0, Eg = 1, Fg = 2, p = null, j = void 0, o = !0, q = !1;

Q = {};

Q.vF = [1.289059486E7, 8362377.87, 5591021, 3481989.83, 1678043.12, 0];
Q.YN = [
    [    1.410526172116255E-8,      8.98305509648872E-6,    -1.9939833816331,   200.9824383106796,  -187.2403703815547,      91.6087516669843, -23.38765649603339, 2.57121317296198, -0.03801003308653, 1.73379812E7],
    [   -7.435856389565537E-9,      8.983055097726239E-6,   -0.78625201886289,  96.32687599759846,  -1.85204757529826,      -59.36935905485877, 47.40033549296737, -16.50741931063887, 2.28786674699375, 1.026014486E7],
    [   -3.030883460898826E-8,      8.98305509983578E-6,     0.30071316287616,  59.74293618442277,   7.357984074871,        -25.38371002664745, 13.45380521110908, -3.29883767235584, 0.32710905363475, 6856817.37],
    [   -1.981981304930552E-8,      8.983055099779535E-6,    0.03278182852591,  40.31678527705744,   0.65659298677277,      -4.44255534477492, 0.85341911805263, 0.12923347998204, -0.04625736007561, 4482777.06],
    [    3.09191371068437E-9,       8.983055096812155E-6,    6.995724062E-5,    23.10934304144901,  -2.3663490511E-4,       -0.6321817810242, -0.00663494467273, 0.03430082397953, -0.00466043876332, 2555164.4],
    [    2.890871144776878E-9,      8.983055095805407E-6,   -3.068298E-8,       7.47137025468032,   -3.53937994E-6,         -0.02145144861037, -1.234426596E-5, 1.0322952773E-4, -3.23890364E-6, 826088.5]
];

Q.rb = function(a){
    if (a === null || a === undefined)
        return new H(0, 0);
    var b, c;
    b = new H(Math.abs(a.lng), Math.abs(a.lat));
    for (var d = 0; d < this.vF.length; d++)
        if (this.vF[d] <= b.lat) {
            c = this.YN[d];
            break
        }
    a = this.hJ(a, c);
    return a = new H(a.lng.toFixed(6), a.lat.toFixed(6))
};

Q.hJ = function(a, b){
    if (a && b) {
        var c = b[0] + b[1] * Math.abs(a.lng),
            d = Math.abs(a.lat) / b[9],
            d = b[2] + b[3] * d + b[4] * d * d + b[5] * d * d * d + b[6] * d * d * d * d + b[7] * d * d * d * d * d + b[8] * d * d * d * d * d * d,
            c = c * (0 > a.lng ? -1 : 1),
            d = d * (0 > a.lat ? -1 : 1);
        return new H(c, d)
    }
};

var lb = function(a, b) {
    if ("string" == typeof a && a) {
        var c = a.split("|"), d, e, f;
        if (1 == c.length)
            d = Gg(a);
        else if (d = Gg(c[2]), e = Gg(c[0]), f = Gg(c[1]), !b)
            return d;
        c = {type: d.aV};
        if (b)
            switch (c.type) {
                case 2:
                    e = new H(d.nd[0][0], d.nd[0][1]);
                    e = Q.rb(e);
                    c.point = e;
                    c.W = [e];
                    break;
                case 1:
                    c.W = [];
                    d = d.nd[0];
                    for (var g = 0, i = d.length - 1; g < i; g += 2) {
                        var k = new H(d[g], d[g + 1]), k = Q.rb(k);
                        c.W.push(k)
                    }
                    e = new H(e.nd[0][0], e.nd[0][1]);
                    f = new H(f.nd[0][0], f.nd[0][1]);
                    e = Q.rb(e);
                    f = Q.rb(f);
                    c.Ua = new eb(e, f)
            }
        return c
    }
};


function Gg(a) {
    var b;
    b = a.charAt(0);
    var c = -1;
    b == "." ? c = 2 : b == "-" ? c = 1 : b == "*" && (c = 0);
    b = c;
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    for (var c = a.substr(1), d = 0, e = c.length, f = [], a = [], g = []; d < e; ) //=vmbNLB2xDjPA;
        if (c.charAt(d) == "=") {
            if (e - d < 13)
                return 0;
            a: {
                for (var g = c.substr(d, 13), i = f, k = 0, l = 0, m = 0, n = 0; n < 6; n++) { //g �? =vmbNLB2xDjPA
                    m = charToNum(g.substr(1 + n, 1));
                    if (0 > m) {
                        g = -1 - n;
                        break a
                    }
                    k += m << 6 * n;
                    m = charToNum(g.substr(7 + n, 1));
                    if (0 > m) {
                        g = -7 - n;
                        break a
                    }
                    l += m << 6 * n
                }
                i.push(k);
                i.push(l);
                g = 0
            }
            if (0 > g)
                return 0;
            d += 13
        } else if (c.charAt(d) == ";")
            a.push(f.slice(0)), f.length = 0, ++d;
        else {
            if (e - d < 8)
                return 0;
            g = Kg(c.substr(d, 8), f);
            if (g < 0)
                return 0;
            d += 8
        }
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    for (c = 0; c < a.length; c++) {
        for (e = 0; e < a[c].length; e++)
            a[c][e] /= 100
    }
    return {aV: b,nd: a}
}

function charToNum(char) { //A-Z a-z 0-9 + / : 0 - 63
    var charNum = char.charCodeAt(0);
    return "A" <= char && char <= "Z" ? charNum - 65 :
                "a" <= char && char <= "z" ? charNum - 71 :
                    "0" <= char && char <= "9" ? charNum + 4 :
                        char == "+" ? 62 : char == "/" ? 63 : -1
}

function Kg(a, numArray) { //将字符串转为-2^23 + 1 ~ 2^23
    var arrayLen = numArray.length;
    if (arrayLen < 2)
        return -1;
    var d = 0, e = 0;
    for (var f = 0, g = 0; g < 4; g++) {
        f = charToNum(a.substr(g, 1));
        if (0 > f)
            return -1 - g;
        d += f << 6 * g;
        f = charToNum(a.substr(4 + g, 1));
        if (0 > f)
            return -5 - g;
        e += f << 6 * g
    }
    d > 8388608 && (d = 8388608 - d);
    e > 8388608 && (e = 8388608 - e);
    numArray.push(numArray[arrayLen - 2] + d);
    numArray.push(numArray[arrayLen - 1] + e);
    return 0
}


function H(a, b) {
    this.lng = a;
    this.lat = b
}

function eb(a, b) {
    a && !b && (b = a);
    a && (this.kl = new H(a.lng, a.lat), this.Uk = new H(b.lng, b.lat), this.ve = a.lng, this.ue = a.lat, this.qe = b.lng, this.pe = b.lat)
}

console.log(lb(".=xDUQLBXyCjPA;", true));
//console.log(lb(".=slbNLBexDjPA;|.=6iTOLBEcjjPA;|-=slbNLBexDjPAPAAAEAAAvWBAbaAAQBEA5+AAT6DAawAA2wBA1VAAMCBAzMAAS7BA2XAAMBAAPAAAwcCAmfAA1UBAGRAAZAAAGAAA3lAAuEAAypAANFAAxKBAUJAA1nAA9EAAPAAACAAA9ZBA7WAAeHBANSAALtCAWiAAvIAAvBAAZtBAsVAAZAAAFAAA9\/AAPIAAPAAACAAAlWAAqAAgNYAAqAAgAAAAAAAAqJAAjRAAyAAAbBAAbNAASvAAgAAAvBAAWGAAYrAAncAAlDDACDAAxUAASAAA5BAAuJAAs\/CAPLAAudDAGAAA1BAABQCAeFAgAAAALAAgYBAACAAgKvAA8AAgb5BAYCAgsdBA1BAgu3BAWCAgCAAAAAAAg2EAWMAgXTCA3FAgvPBALDAgXHDA7HAgAAAAAAAAyVAAfJAAqAAASAAAcNAAbSEAZCAAzwAALAAgAAAABAAAWAAAFGAAOFCA2AAAvSAAvIAAv\/CAuBAAmlAAEAAA1AAAMNAAKhCAJEAAoyAAGmAAOVAAnWAAYAAAZWAA0UAgHUAAsSAgPiAA0fAgKwAAwsAgpTAAFhAgj4AANfBgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA;", true));
