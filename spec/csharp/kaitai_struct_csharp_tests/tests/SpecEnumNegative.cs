using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecEnumNegative : CommonSpec
    {
        [Test]
        public void TestEnumNegative()
        {
            var r = EnumNegative.FromFile(SourceFile("enum_negative.bin"));
            Assert.AreEqual(r.F1, EnumNegative.Constants.NegativeOne);
            Assert.AreEqual(r.F2, EnumNegative.Constants.PositiveOne);
        }
    }
}
