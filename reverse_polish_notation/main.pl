use strict;
use warnings;
use utf8;
use Carp;

# 逆ポーランド記法を計算する
sub calculate {
    my $formula = shift;
    print $formula . ': ';

    # 式を分割
    my @f = split (/ /, $formula);
    my @stuck;
    my $r = 0;
    my $l = 0;

    for my $v (@f) {
        # 数値ならスタックに積む
        if ($v =~ /^\d$/) {
            push @stuck, int($v);
        # 式ならスタックから数値を取得して計算処理後スタックに積む
        } elsif ($v eq '+') {
            $r = pop @stuck;
            $l = pop @stuck;
            push @stuck, ($l + $r);
        } elsif ($v eq '-') {
            $r = pop @stuck;
            $l = pop @stuck;
            push @stuck, ($l - $r);
        } elsif ($v eq '*') {
            $r = pop @stuck;
            $l = pop @stuck;
            push @stuck, ($l * $r);
        } elsif ($v eq '/') {
            $r = pop @stuck;
            $l = pop @stuck;
            push @stuck, ($l / $r);
        # それ以外はエラー
        } else {
            croak 'formula is wrong format!';
        }
    }
    croak 'failed calculate!' unless (scalar @stuck == 1);

    return pop @stuck;
}


# formula='3 4 + 1 2 - *': -7
print calculate('3 4 + 1 2 - *') . "\n";

# formula='1 2 3 * +': 7
print calculate('1 2 3 * +') . "\n";
