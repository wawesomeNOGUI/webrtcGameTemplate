function smoothstep (start, fin, t)
{
   if (t <= 0)
      return start;

   if (t >= 1)
      return fin;

    var dis = fin - start;

    // derived smoothstep = -2x^3 + 3x^2
   return start + dis * t*t*(3 - 2*t);
}