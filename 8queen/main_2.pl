use strict;
use warnings;
use utf8;
use Carp;
use Data::Dumper;

# 8クイーン
#「8×8のチェス盤上にクイーンを8個置き， どのクイーンからも他のクイーンの場所に1手では行けないような配置にせよ」
#
# バックトラック法
# 解の候補をすべて調べることを組織的にかつ効率よく行うための技法
# 8個のクイーンをチェス盤上に置く置き方が 64C8 = 4426165368 通り
# しかし1行に1つしかおけないから、1行につき8パターンで8行ある 8*8*8*8*8*8*8*8 = 16777216 通り

# 再帰関数チェックする方式
our $NUM = 8;

sub print_queen {
    my $queen = shift;

    for (my $i = 0; $i < $NUM; $i++) {
        printf("%d ", $queen->[$i]);
    }
    print("\n");
}

sub check {
    my $queen = shift;

    for (my $i = 0; $i < $NUM-1; $i++) {
        for (my $j = $i+1; $j < $NUM; $j++) {
            return 0 if ($queen->[$i] == $queen->[$j] || abs($queen->[$i] - $queen->[$j]) == $j - $i);
        }
    }
    return 1;
}

sub set_queen {
    my ($queen, $i) = @_;

    if ($i == $NUM) {
        print_queen($queen) if (check($queen));
        return;
    }

    for (my $j = 0; $j < $NUM; $j++) {
        $queen->[$i] = $j;
        set_queen($queen, $i+1);
    }
}

sub main {
    set_queen([0 .. $NUM-1], 0);
}
main();
