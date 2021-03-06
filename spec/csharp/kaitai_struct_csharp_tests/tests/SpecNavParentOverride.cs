using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecNavParentOverride : CommonSpec
    {
        [Test]
        public void TestNavParentOverride()
        {
            var r = NavParentOverride.FromFile(SourceFile("nav_parent_codes.bin"));
            Assert.AreEqual(r.ChildSize, 3);
            Assert.AreEqual(r.Child1.Data, new byte[] { 73, 49, 50 });
            Assert.AreEqual(r.Mediator2.Child2.Data, new byte[] { 51, 66, 98 });
        }
    }
}
