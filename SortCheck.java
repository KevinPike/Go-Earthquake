// Check whether a file is sorted.
// Records are assumed to be 4 bytes with a 2-byte key and 2-byte data field
import java.io.*;
import java.nio.*;

public class SortCheck {

static final int NumRec = 1024;
static final int RecSize = 4;

public static void main(String args[]) throws IOException {

  if (args.length != 1) {
    System.out.println("Usage: SortCheck <file>");
    return;
  }

  File f = new File(args[0]);
  int filelength = (int)f.length();
  if (((filelength/(RecSize*NumRec))*(RecSize*NumRec)) != filelength) {
    System.out.println("The file size must be a multiple of " + RecSize*NumRec);
    return;
  }
  int blocks = filelength/(RecSize * NumRec);

  DataInputStream theFile = new DataInputStream(
          new BufferedInputStream(new FileInputStream(args[0])));

  byte[] rec1b = new byte[RecSize];
  ByteBuffer rec1bb = ByteBuffer.wrap(rec1b);
  ShortBuffer rec1sb = rec1bb.asShortBuffer();

  byte[] rec2b = new byte[RecSize];
  ByteBuffer rec2bb = ByteBuffer.wrap(rec2b);
  ShortBuffer rec2sb = rec2bb.asShortBuffer();

  int i;
  theFile.read(rec2b, 0, 4); // Priming read
  for (i=1; i<filelength/RecSize; i++) {
    // Along with checking the correctness,
    // we also output the value at the start of each 4096-byte block,
    // where the records are interpreted as a short int key and short int data value
    if(((i-1)%1024)==0) System.out.format("%5d %5d  ", rec2sb.get(0), rec2sb.get(1));
    if(((i)%8192)==0) System.out.println();
    System.arraycopy(rec2b, 0, rec1b, 0, RecSize);

    theFile.read(rec2b, 0, 4); // Priming read
    if (rec1sb.get(0) > rec2sb.get(0)) {
      System.out.println("Record " + i + ": Unsorted output file");
      return;
    }
  }
  System.out.println();
  theFile.close();
  System.out.println(i + " records processed");
}

}