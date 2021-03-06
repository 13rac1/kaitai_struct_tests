package spec::perl::TestNestedTypes;

use strict;
use warnings;
use base qw(Test::Class);
use Test::More;
use NestedTypes;

sub test_nested_types: Test(3) {
    my $r = NestedTypes->from_file('src/fixed_struct.bin');

    is($r->one()->typed_at_root()->value_b(), 80, 'Equals');
    is($r->one()->typed_here()->value_c(), 65, 'Equals');
    is($r->two()->value_b(), 67, 'Equals');
}

Test::Class->runtests;
