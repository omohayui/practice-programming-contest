use strict;
use warnings;
use utf8;
use Carp;

# 1以上の整数nが与えられたとき
# 1からnまでの整数のうち
# 5で割り切れない数の合計

# 回答1
sub count {
    my $n = shift;
    # 引数チェック
    croak 'Not integer!' unless (($n ^ $n) eq '0');
    croak 'Not positive integer!' unless ($n > 0);

    my $total = 0;
    for (my $i = 1; $i <= $n; $i++) {
        $total += $i if ($i % 5) > 0;
    }
    return $total;
}

# 回答2
sub count2 {
    my $n = shift;
    # 引数チェック
    croak 'Not integer!' unless (($n ^ $n) eq '0');
    croak 'Not positive integer!' unless ($n > 0);

    # 1からnまでの整数の合計 - 5の倍数の合計
    return $n*($n+1)/2 - ($n/5)*($n/5+1)*5/2;
}

# n=5: 10
print count(5) . "\n";
print count2(5) . "\n";

# n=100: 4000
print count(100) . "\n";
print count2(100) . "\n";

# n="あ": error
print count("あ") . "\n";

# n=-5: error
print count(-5) . "\n";
