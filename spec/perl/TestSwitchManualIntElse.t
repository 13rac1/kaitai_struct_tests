package spec::perl::TestSwitchManualIntElse;

use strict;
use warnings;
use base qw(Test::Class);
use Test::More;
use SwitchManualIntElse;

sub test_switch_manual_int_else: Test(9) {
    my $r = SwitchManualIntElse->from_file('src/switch_opcodes2.bin');

    is(scalar(@{$r->opcodes()}), 4, 'Equals');

    is($r->opcodes()->[0]->code(), 83, 'Equals');
    is($r->opcodes()->[0]->body()->value(), 'foo', 'Equals');

    is($r->opcodes()->[1]->code(), 88, 'Equals');
    is($r->opcodes()->[1]->body()->filler(), 0x42, 'Equals');

    is($r->opcodes()->[2]->code(), 89, 'Equals');
    is($r->opcodes()->[2]->body()->filler(), 0xcafe, 'Equals');

    is($r->opcodes()->[3]->code(), 73, 'Equals');
    is($r->opcodes()->[3]->body()->value(), 7, 'Equals');
}

Test::Class->runtests;
