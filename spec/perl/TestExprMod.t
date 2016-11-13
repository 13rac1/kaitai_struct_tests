package spec::perl::TestExprMod;

use strict;
use warnings;
use base qw(Test::Class);
use Test::More;
use ExprMod;

sub test_expr_mod: Test(6) {
    my $r = ExprMod->from_file('src/fixed_struct.bin');

    is($r->int_u(), 1262698832, 'Equals');
    is($r->int_s(), -52947, 'Equals');

    is($r->mod_pos_const(), 9, 'Equals');
    is($r->mod_neg_const(), 4, 'Equals');
    is($r->mod_pos_seq(), 5, 'Equals');
    is($r->mod_neg_seq(), 2, 'Equals');
}

Test::Class->runtests;
