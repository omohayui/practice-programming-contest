use strict;
use warnings;
use utf8;
use Carp;
use Data::Dumper;

# From https://leetcode.com/problems/two-sum/
# 整数の配列が与えられたとき、それらが特定のターゲットになるように2つの数のインデックスを返します。
# 各入力は厳密に1つの解を持つと仮定することができ、同じ要素を2回使用することはできません。
sub two_sum {
    my ($nums, $target) = @_;

    # 引数チェック
    croak 'target is not integer!' unless (($target ^ $target) eq '0');
    croak 'nums is not array!' unless (ref($nums) eq 'ARRAY');

    # nums の要素数
    my $len = scalar @$nums;
    for (my $i = 0; $i < $len; $i++) {
        for (my $j = 0; $j < $len; $j++) {
            if ($i != $j && $nums->[$i] + $nums->[$j] == $target) {
                return [$i, $j];
            }
        }
    }
    # 1つも解がない場合はエラー
    croak 'failed two_sum!';
}


# Given nums = [2, 7, 11, 15], target = 9
# return [0, 1]
print Dumper(two_sum([2, 7, 11, 15], 9));

# Given nums = [2, 7, 11, 15, 20], target = 35
# return [3, 4]
print Dumper(two_sum([2, 7, 11, 15, 20], 35));
