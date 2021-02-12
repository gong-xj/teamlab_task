
import SwiftUI

struct ScView: View {
    @State var scData: [Sc]
    @State var name: String
    
    
    var body: some View {
        VStack(alignment: .center){
            Text(name)
            List(scData) { sc in
                Text(sc.nameAndScore)
            }
        }
    }
}

struct ScView_Previews: PreviewProvider {
    static var previews: some View {
        ScView(scData: [], name: "")
    }
}
